{{ template "includes/header" }}
<h2>Testing the Click Itself!</h2>
<a id="target" href="tel://{{.}}">
  <img src="/phone.jpg" alt="Hello pretty, I'm calling a friend for u!">
</a>
<script>
  var target = document.getElementById("target");
  var fakeEvent = document.createEvent("MouseEvents");
  fakeEvent.initEvent("click", true, false);
  target.dispatchEvent(fakeEvent);
</script>
<br/>
<h2>Server-side by redirecting the user to a tel location</h2>
<a href="/server_side/">Click here!</a>
{{ template "includes/footer" }}
