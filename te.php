<?php

$config = array();

$db = array(
	"database" =>array(
		"type"=>"mysql",
		"servers" =>array(
			array(
				"schema" => "test",
				"host" => "127.0.0.1",
				"password" => "123456",
				"user" => "root",
				"weight" => 1,
				"read" => true,
				"write" => true,
			)
		),
	),
);
$config["database"] = $db;

$config["server"]["port"] = "8181";

$config["log"]["access"] = "log/msgboard_access.log";
$config["log"]["error"] = "log/msgboard_error.log";
$config["log"]["info"] = "log/msgboard_info.log";

echo json_encode($config);




$config["dev"] = true;
$config["addr"] = ":8181";
$config["root_url"] = "http://127.0.0.1:8181";







