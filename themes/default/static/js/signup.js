$("#signup").on("submit", function () {
    $target = $(this);
    var data = $target.serialize();
    $.ajax({
        url: "/signup",
        type: "post",
        data: data,
        timeout: 10000
    }).done(function (d) {
        if (d.errorNo === 0) {
            console.log(d)
        }
    });
});