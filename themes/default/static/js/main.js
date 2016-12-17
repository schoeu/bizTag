$addsForm = $("#addsForm");
$smtBtn = $addsForm.find('.addsBtn');
$smtBtn.on("click", function () {
    $smtBtn.attr("disabled", true);
    var data = $addsForm.serialize();
    $.ajax({
        url: "/addsite",
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