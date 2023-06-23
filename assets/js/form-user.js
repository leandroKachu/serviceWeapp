$('#form_user').on('submit', createAccount);


function createAccount(event) {

    event.preventDefault();

    if ($('#password').val() != $('#confirm_password').val()) {
        alert("Password is diferent");
        console.log($('#password').val(), $('#confirm_password').val())
        return;
    }

}