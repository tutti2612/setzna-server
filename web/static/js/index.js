// ローカルで動かすときはWebsocketプロトコルをws、Herokuで動かすときはwssに変更してください。
let url = "ws://" + window.location.host + window.location.pathname + "/ws?latitude=66.12312&longitude=123.34234";
let ws = new WebSocket(url);
let name = localStorage.getItem("user")
let chat = document.getElementById("chat");

const text = document.getElementById("text");

ws.onmessage = function (msg) {
  let obj = JSON.parse(msg.data);
  obj.conten = escape_html(obj.content);
  let line ="";
  if (obj.name==name){
    line =`<div class='line-right'>
            <p class='line-right-text'>${obj.content} </p>
            <div class="line-right-time">${now()}</div>
           </div>`
  }else{
    let image = '<img src="/static/img/icon.png"/>'
    line =`<div class='line-left'>
                ${image}
                <div class='line-left-container'>
                    <p class='line-left-name'>
                    ${obj.name}
                    </p>
                    <p class='line-left-text'>
                    ${obj.content}
                    </p>
                    <div class='line-left-time'>
                        ${now()}
                    </div>
                </div>
           </div>`
  }
  chat.innerHTML += line;
};

text.onkeydown = function (e) {
  if (e.keyCode === 13) {
    send_data();
  }
};

function send_data(){
    if (text.value == "")return;
    text.value = escape_html(text.value);
    let sendData = `{"name":"${name}","content":"${text.value}","latitude":"65.123123","longitude":"123.123123","type":"post"}`;
    ws.send(sendData);
    text.value = "";
}

function now() {
    let date = new Date();
    let min = (date.getMinutes()<10)?`0${date.getMinutes()}`:date.getMinutes();
    let hour = (date.getHours()<10)?`0${date.getHours()}`:date.getHours();
    return `${hour}:${min}`
};

function escape_html (string) {
    if(typeof string !== 'string') {
      return string;
    }
    return string.replace(/[&'`"<>]/g, function(match) {
      return {
        '&': '&amp;',
        "'": '&#x27;',
        '`': '&#x60;',
        '"': '&quot;',
        '<': '&lt;',
        '>': '&gt;',
      }[match]
    });
}
