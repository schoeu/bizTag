$suForm = $("#signup");
$smtBtn = $suForm.find('.signupSbm');
$suForm.on("submit", function () {
    $smtBtn.attr("disabled", true);
    var data = $target.serialize();
    $.ajax({
        url: "/signup",
        type: "post",
        data: data
    }).done(function (d) {
        if (d.errorNo === 0) {
            console.log(d)
        }
    }).complete(function () {
        $smtBtn.attr("disabled", false);
    });
});