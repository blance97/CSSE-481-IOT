$(function() {
    var request = $.ajax({
        type: 'GET',
        url: '/listFile',
        async: false,
        success: function(data) {
            var obj = jQuery.parseJSON(data)
            for (i = 0; i < obj.Rooms.length; i++) {
                $('#files').append($('<option>', {
                    value: obj.Rooms[i],
                    text: obj.Rooms[i]
                }));
            }
        }
    });

});


function loadData1() {
      var datas = $("#files").val()
    var request = $.ajax({
        type: 'GET',
        url: '/getFile/?data=' + datas,
        async: false,
        success: function(data) {
          console.log("Changed data");

        }
    });
}
