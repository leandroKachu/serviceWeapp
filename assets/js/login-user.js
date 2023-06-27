$('#login').on('submit', logUser);


function logUser(event) {

    event.preventDefault();
    $.ajax({
        url: apiPath(),
        method: 'POST',
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        alert("pass or email is wrong")
    });


    function apiPath() {
        return "/login";
    }
}