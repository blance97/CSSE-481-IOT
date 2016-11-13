$(function() {
    var table = '<table>' +
        '<tr>' +
        '<th>Time</th>' +
        '<th>Temp</th>' +
        '<th>Abs Pressure</th>' +
        '<th>Rel Pressure</th>' +
        '<th>Humidity</th>' +
        '<th>Rain Sensor 1</th>' +
        '<th>Rain Sensor 2</th>' +
        '</tr>'
    var request = $.ajax({
        type: 'GET',
        url: '/getData/?data=Temp',
        async: false,
        success: function(data) {
            var obj = jQuery.parseJSON(data)
            if (obj == null) {
                return
            }
            var array;
            for (var i = 0; i < obj.length; i++) {
               var index = i+1;
                table += '<tr><td>' +
                    index + '</td><td>' +
                    obj[i].Temp + '</td><td>' +
                    obj[i].Absolute_Pressure + '</td><td>' +
                    obj[i].Relative_Pressure + '</td><td>' +
                    obj[i].Humidity + '</td><td>' +
                    obj[i].Rain_Sensor_1 + '</td><td>' +
                    obj[i].Rain_Sensor_2 + '</td></tr>'
                table += '</table>'
            }
            document.getElementById('hi').innerHTML = table
        }
    });
});
