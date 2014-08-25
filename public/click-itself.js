var target = document.getElementById("target");
var fakeEvent = document.createEvent("MouseEvents");
fakeEvent.initEvent("click", true, false);
target.dispatchEvent(fakeEvent);
