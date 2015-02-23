$(function() {
  var conn;
  var reconnectIntervalId;

  var connect = function(){
    conn = new WebSocket("ws://localhost:8080/ws");
    return conn.readyState === 1;
  }

  var reconnect = function() {
    reconnectIntervalId = setInterval(function(){
      if(conn.readyState !== 1){
        connect();
      } else {
        clearInterval(reconnectIntervalId);
      }
    }, 5000);
  }

  if (window["speechSynthesis"] && window["WebSocket"]) {
    if (!connect()){
      reconnect();
    }

    conn.onclose = function(e) {
      reconnect();
    }

    conn.onmessage = function(e) {
      if($("#speaker").is(":checked")) {
        var data = JSON.parse(e.data);
        var u = new SpeechSynthesisUtterance();
        u.text = data.text;
        u.lang = data.lang;
        window.speechSynthesis.speak(u);
      }
    }

    var VALID_LANGS = {
      "de": "de-DE",
      "da": "da-DK",
      "en": "en-US",
      "gb": "en-GB",
      "es": "es-ES",
      "fi": "fi-FI",
      "fr": "fr-FR",
      "no": "nb-NO",
      "ru": "ru-RU",
      "sv": "sv-SE"
    };

    $("#text-field").focus();
    $("#text-field")[0].setSelectionRange(3, 3);

    $("#text-field").keydown(function(e) {
      if (e.keyCode == 13) {
        var text = $.trim($("#text-field").val());
        var prefix = text.substr(0,text.indexOf(":") + 1);
        var data = {};

        var lang = prefix.substr(0,prefix.indexOf(":"));
        if (Object.keys(VALID_LANGS).indexOf(lang) != -1) {
          data.lang = VALID_LANGS[lang];
          data.text = $.trim(text.substr(lang.length + 1, text.length));

        } else {
          lang = "sv";
          data.lang = "sv-SE";
          data.text = text;
        }

        $("#text-field").val(lang + ":");
        conn.send(JSON.stringify(data));

        return false;
       }
    });

    $("#flash").hide();

  } else {
    $("#not-supported").show();
  }

});
