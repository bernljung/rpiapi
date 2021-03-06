$(function() {
  var conn;
  var reconnectIntervalId;
  var reconnecting = false;

  var connect = function(){
    console.log("Connecting to WebSocket.");
    conn = new WebSocket("ws://" + window.location.host + "/ws");
    bindSocketEvents();
  }

  var reconnect = function() {
    reconnecting = true;
    console.log("Reconnecting.");
    connect();
    reconnectIntervalId = setInterval(function(){
      if(conn.readyState !== 1){
        connect();
        if(conn.readyState === 1){
          console.log("Reconnected.");
          reconnecting = false;
        }
      } else {
        console.log("Reconnected.");
        reconnecting = false;
        clearInterval(reconnectIntervalId);
      }
    }, 5000);
  }

  var bindSocketEvents = function() {
    conn.onopen = function(e) {
      console.log("Connection opened.");
    }

    conn.onclose = function(e) {
      console.log("Connection closed.");
      if(!reconnecting){
        reconnect();
      }
    }

    conn.onerror = function(e) {
      console.log("Connection error.");
      if(!reconnecting){
        reconnect();
      }
    }

    conn.onmessage = function(e) {
      if(window["speechSynthesis"] && $("#speaker").is(":checked")) {
        var data = JSON.parse(e.data);
        var u = new SpeechSynthesisUtterance();
        u.text = data.text;
        u.lang = data.lang;
        window.speechSynthesis.speak(u);
      }
    }
  }

  if (window["WebSocket"]) {
    connect();

    bindSocketEvents();

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

    if(!window["speechSynthesis"]) {
      $("#no-speak").show();
      $("#speaker").hide();
    }

  } else {
    $("#not-supported").show();
  }

});
