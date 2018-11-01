<?php
	//  This is a C&C server for a fake botnet. DCI CTF 2018: https://www.facebook.com/events/461004354417960/ 

	function crever()
	{
		sleep(5);
		http_response_code(404);
		header($_SERVER["SERVER_PROTOCOL"]." 404 Not Found", true, 404);
		readfile('error.htm');
		echo("<!-- 8pbnyk93pkJ6l/m4PyJJ/sYmhM4+eT5+cSyEkuHiKbQ= -->");
		die();
	}

	function showflag()
	{
		sleep(2);
		http_response_code(404);
		header($_SERVER["SERVER_PROTOCOL"]." 404 Not Found", true, 404);
		echo(base64_encode("DCI{Linux.Trojan.Not.A.Keylogger}"));
		die();
	}

	if(isset($_GET['dog']) && isset($_COOKIE['maga']) && isset($_SERVER['CONTENT_LENGTH']) && isset($_SERVER['HTTP_USER_AGENT']))
	{
		usleep(500 * 1000);

		date_default_timezone_set('US/Eastern');
		$v_ip = $_SERVER['REMOTE_ADDR'];
		$v_date = date("l d F H:i:s");
		$v_agent = $_SERVER['HTTP_USER_AGENT'];
		$v_param = $_GET['dog'];
		$v_uid = $_COOKIE['maga'];	// Les cookies rendent l'Amérique encore meilleure
		$v_post = file_get_contents('php://input');
		$v_len = $_SERVER['CONTENT_LENGTH'];
		$v_mesuredlen = strlen($v_post);

		if (strlen($v_param) >= 6 && strlen($v_param) <= 45 && strlen($v_uid) == 32 && $v_agent == "curl/7.37.0")
		{
			if ($v_len == $v_mesuredlen && $v_mesuredlen == 20)
			{
				$fp = fopen("amigos/bots.txt", "a");
				fputs($fp, "UID: $v_uid - IP: $v_ip - DATE: $v_date - BROWSER: $v_agent - Param: $v_param - Post: $v_post\r\n");
				fclose($fp);

				showflag();
			}
		}
	}
	else
	{
		// Tentatives d'accès à la page directement
		$v_ip = $_SERVER['REMOTE_ADDR'];
		$v_date = date("l d F H:i:s");
		$v_agent = $_SERVER['HTTP_USER_AGENT'];

		$fp = fopen("amigos/tent.txt", "a");
		fputs($fp, "IP: $v_ip - DATE: $v_date - BROWSER: $v_agent\r\n");
		fclose($fp);
	}

	crever();
?>
