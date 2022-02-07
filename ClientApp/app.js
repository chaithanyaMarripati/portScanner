function SearchClicked() {
let socket = new WebSocket("wss:url to be inserted");

// send message from the form
document.forms.publish.onsubmit = function() {
  let outgoingMessage = document.getElementById("inputid").value;

  socket.send(outgoingMessage);
  return false;
};

// message received - show the message in div#messages
socket.onmessage = function(event) {
  let message = event.data;

  let messageElem = document.createElement('div');
  messageElem.textContent = message;
  document.getElementById('messages').prepend(messageElem);
}
}