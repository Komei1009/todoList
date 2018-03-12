// ページ読み込み時に実行したい処理
$(document).ready( function(){
    var mode = sessionStorage.getItem('mode')
    if (mode == 0 || mode == null){
        display("all")
    }else if (mode == 1){
        display("active")
    }else{
        display("completed")
    }

});

// タスク追加リクエスト
function addTodo(){
    var data = {'task': $('#inputTask').val()};

    $.ajax({
        url: '/addTodo',
        type:'POST',
        dataType: 'json',
        data : data,
        headers: {
            'Authorization': localStorage.getItem('token')
        },
        timeout:10000,
    }).done(function (result) {
        console.log(result.ok)
    }).fail(function (XMLHttpRequest, textStatus, errorThrown) {
        $errMsg = XMLHttpRequest['responseJSON']['err']
        alert($errMsg);
    })
}

// タスクステータス変更リクエスト
function todoControll(){
    $("#todos td").off().bind('click', function(){
        var task_name
        $tag_td = $(this)[0]
        $tag_tr = $(this).parent()[0]
        console.log("%s行, %s列", $tag_tr.rowIndex, $tag_td.cellIndex);

        if ($tag_td.cellIndex == "3"){
            // タスク名取得
            task_name = $(this).prev().prev().prev().text()
            console.log(task_name)
            controllStatus(task_name,"change")
        }else if($tag_td.cellIndex == "4"){
            // タスク名取得
            var task_name = $(this).prev().prev().prev().prev().text()
            console.log(task_name)
            controllStatus(task_name,"delete")
        }

        location.reload(true)
        return
    });
}

// ステータス変更
function controllStatus(task_name,controll){
    var data = {task: task_name, controll : controll};
    $.ajax({
        url: '/controllerTodo',
        type:'POST',
        dataType: 'json',
        data : data,
        headers: {
            'Authorization': localStorage.getItem('token')
        },
        timeout:10000,
    }).done(function (result) {
        console.log(result.ok)
    }).fail(function (XMLHttpRequest, textStatus, errorThrown) {
        $errMsg = XMLHttpRequest['responseJSON']['err']
        alert($errMsg);
    })
}

// 表示テーブル
function display(select){
    var data = {'display' : select}
    $.ajax({
        url: '/todoDisplay',
        type:'POST',
        dataType: 'json',
        data : data,
        headers: {
            'Authorization': localStorage.getItem('token')
        },
        timeout:10000,
    }).done(function(data){
        for (var i = 0; i < data['todos'].length; i++){
            var task_name = data['todos'][i].task_name
            var record_date = new Date(data['todos'][i].CreatedAt)
            // テーブルの要素取得
            var table = document.getElementById("todos")
            // 行を行末に追加
            var row = table.insertRow(-1)
            //td分追加
            var cell1 = row.insertCell(-1)
            var cell2 = row.insertCell(-1)
            var cell3 = row.insertCell(-1)
            var cell4 = row.insertCell(-1)
            var cell5 = row.insertCell(-1)

            // セルの内容入力
            cell1.innerHTML = task_name
            cell2.innerHTML = record_date.getFullYear() + "/" + record_date.getMonth() + "/" + record_date.getDate()+ " " + record_date.getHours() + ":" + record_date.getMinutes()
            cell3.innerHTML = data['todos'][i].status
            cell4.innerHTML = '<button id = "controllButton" onClick= "todoControll()" class="btn btn-default">Change</button>'
            cell5.innerHTML = '<button id = "addButton" onClick="todoControll()" class="btn btn-default">Delete</button>'
        }
    }).fail(function(XMLHttpRequest, textStatus, errorThrown) {
        $errMsg = XMLHttpRequest['responseJSON']['err']
        alert($errMsg);
    })
}

// 表示選択
function displaySelect(ele){
    if (ele.id == "allButton"){
        sessionStorage.setItem('mode',0)
    }else if (ele.id == "activeButton"){
        sessionStorage.setItem('mode',1)
    }else{
        sessionStorage.setItem('mode',2)
    }
    location.reload(true)
}

