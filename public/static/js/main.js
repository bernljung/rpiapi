window.onload = function () {
  $('#speech').focus();
  $('#speech')[0].setSelectionRange(2, 2);
  $('#speech').keydown(function(e) {
    if (e.keyCode == 13) {
      var text = $.trim($('#speech').val());
      var prefix = text.substr(0,2);

      switch(prefix) {
        case "s:":
          $('#lang').val("sv");
          $('#speech').val(text.substr(2, text.length));
          break;
        case "e:":
          $('#lang').val("en");
          $('#speech').val(text.substr(2, text.length));
          break;
        default:
          $('#lang').val("sv");
          break;
      }

      $(this.form).submit();

      if ($('#lang').val() === "en") {
        $('#speech').val("e:");
      } else {
        $('#speech').val("s:");
      }
      return false;
     }
  });

  $('#flash.alert-success').fadeOut(7000);
};