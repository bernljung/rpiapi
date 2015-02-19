window.onload = function () {
  $('#speech').focus();
  $('#speech').keydown(function(e) {
    if (e.keyCode == 13) {
        $(this.form).submit()
        return false;
     }
  });
  $('#speech').keyup(function(e) {
    var text = $('#speech').val();
    if(text.length == 2){
      switch (text){
        case "s:":
          $('#lang option[value=sv]').prop('selected', 'selected');
          $('#lang option[value=en]').prop('selected', null);
          $('#speech').val("");
          break;
        case "e:":
          $('#lang option[value=en]').prop('selected','selected');
          $('#lang option[value=sv]').prop('selected', null);
          $('#speech').val("");
          break;
        default:
          break;
      }
    }
  });
  $('#flash.alert-success').fadeOut(7000);
};