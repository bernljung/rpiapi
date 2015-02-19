window.onload = function () {
  $('#speech').focus();
  $('#speech')[0].setSelectionRange(2, 2);
  $('#speech').keydown(function(e) {
    if (e.keyCode == 13) {
      var text = $('#speech').val();
      var prefix = text.substr(0,2);

      if(prefix === "e:" || prefix == "s:"){
        $('#speech').val(text.substr(2, text.length));
      }
      $(this.form).submit();

      switch($('#lang').val()){
        case "en":
          $('#speech').val("e:");
          break;
        case "sv":
          $('#speech').val("s:");
          break;
        default:
          break
      }
      return false;
     }
  });

  $('#speech').keyup(function(e) {
    var text = $('#speech').val();
    if(text.length == 2){
      switch (text){
        case "s:":
          $('#lang').prop('value', 'sv');
          break;
        case "e:":
          $('#lang').prop('value', 'en');
          break;
        default:
          break;
      }
    }
  });
  $('#flash.alert-success').fadeOut(7000);
};