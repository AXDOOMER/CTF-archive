<!-- Fake Login Panel for a fake Botnet. This is the homepage of the C&C server. DCI CTF 2018: https://www.facebook.com/events/461004354417960/ -->


<!DOCTYPE html>
<html lang="zxx"><head>
<meta http-equiv="content-type" content="text/html; charset=UTF-8">
	<title>Login Panel</title>
	<!-- Meta-Tags -->
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta charset="utf-8">
	<meta name="keywords" content="Login Panel Fake Botnet">
	<script>
		addEventListener("load", function () {
			setTimeout(hideURLbar, 0);
		}, false);

		function hideURLbar() {
			window.scrollTo(0, 1);
		}
	</script>
	<!-- //Meta-Tags -->
	<!-- Stylesheets -->
	<link href="style.css" rel="stylesheet" type="text/css">
	<!--// Stylesheets -->
	<!-- online fonts -->
	<link href="css.css" rel="stylesheet">
</head>

<body>

<!---728x90--->

	<header>
		<h1 class="text-center">Login Panel</h1>
	</header>
	<!---728x90--->

	<div class="w3ls-login">
		<p class="subcsribe-text">Enter your login information</p>
		<form class="subscribe-form" action="" method="post">
			<div class="agile-field-txt">
				<input id="username" name="Name" placeholder=" " required="" type="text">
				<label for="username">Name
					<span>*</span>
				</label>
			</div>
			<div class="agile-field-txt">
				<input id="password" name="Password" placeholder=" " required="" type="password">
				<label for="password">Password
					<span>*</span>
				</label>
			</div>
			<div class="agile-field-txt">
				<button type="submit" class="btn" value="Submit"> Login</button>
			</div>
		</form>
	</div>
	<!---728x90--->

	<!--copyright-->
	<div class="copy-wthree">
		<p>© 2018 The Russian Mafia. All Rights Reserved | Design by
			<a href="https://www.google.com/search?q=Lycanthrope" target="_blank">RedPill</a>
		</p>
	</div>
	<!--//copyright-->

<?php
	usleep(500 * 1000);

	if(isset($_POST['Name']) && isset($_POST['Password']))
	{
		date_default_timezone_set('US/Eastern');
		$v_ip = $_SERVER['REMOTE_ADDR'];
		$v_date = date("l d F H:i:s");
		$v_agent = $_SERVER['HTTP_USER_AGENT'];
		$v_pass = $_POST['Password'];
		$v_user = $_POST['Name'];

		$fp = fopen("amigos/logins.txt", "a");
		fputs($fp, "IP: $v_ip - DATE: $v_date - BROWSER: $v_agent - User: $v_user -  Pass: $v_pass\r\n");
		fclose($fp);
	}
?>

</body></html>
