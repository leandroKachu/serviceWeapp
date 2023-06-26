$('#form_user').on('submit', createAccount);


function createAccount(event) {

    event.preventDefault();

    if ($('#password').val() != $('#confirm_password').val()) {
        alert("Password is diferent");
        console.log($('#password').val(), $('#confirm_password').val())
        return;
    }

    $.ajax({
        url: apiPath(),
        method: 'POST',
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val(),
        }
    });


    function apiPath() {
        return "/send-api";
    }
}