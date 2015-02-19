window.onload = function () {
  $('#speech').focus();
  $('#speech').keydown(function(e) {
    if (e.keyCode == 13) {
        $(this.form).submit()
        return false;
     }
  });
  $('#flash.alert-success').fadeOut(7000);
};