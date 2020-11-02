<?php
ini_set('error_reporting', E_ALL);
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);

include_once("rb.php");
$password="0nm5b1ju";
if(isset($_SERVER["HTTP_PASSWORD"]) and $_SERVER["HTTP_PASSWORD"]==$password){
    $mode="none";
    R::setup( 'mysql:host=localhost;dbname=syslogManager','syslog', 'password' ); 

    if(isset($_SERVER['HTTP_MODE'])){$mode=$_SERVER['HTTP_MODE'];}
    
    /**************************
     * 
     * R::store( $book );
		Raw          string `json:"raw"`
		Priority     int8   `json:"priority"`
		PriorityName string `json:"priorityName"`
		Subject      int8   `json:"subjectData"`
		Topic        string `json:"topics"`
		Date         string `json:"date"`
		Timestamp    int64  `json:"timestamp"`
		TimeUTC      string `json:"timeutc"`
		TimeNow      int64  `json:"timenow"`
		Msg          string `json:"msg"`
		TypeOf       int8   `json:"typeOf"` //1=RFC5424 2=RFC3164 0=RAW
		TypeOfName   string `json:"typeOfName"`
     */
    if($mode=="PUT_NEW_ROW"){
     
    $json = file_get_contents('php://input');
    $obj = json_decode($json);
    var_dump($obj->raw);
   
    $book = R::dispense( 'mainrall' );
    $book->raw=$obj->raw;
    $book->priority=$obj->priority;
    $book->subjectData=$obj->subjectData;
    $book->topics=$obj->topics;
    $book->date=$obj->date;
    $book->timestamp=$obj->timestamp;
    $book->timeutc=$obj->timeutc;
    $book->timenow=$obj->timenow;
    $book->msg=$obj->msg;
    $book->TypeOf=$obj->TypeOf;
    $book->typeOfName=$obj->typeOfName;
   
    R::store( $book );
    //$book->raw = $obj['rawData'];
    /*
    $book->priority = $obj['priority'];
    $book->priorityName = $obj['priorityName'];
    $book->subject = $obj['subjectData'];
     /*
    $book->topics = $obj['topics'];
    $book->date = $obj['date'];
    $book->timestamp = $obj['timestamp'];
    $book->timeUTC = $obj['timeutc'];
    $book->timeNow = $obj['timenow'];
    $book->msg = $obj['msg'];
    $book->typeOf = $obj['typeOf'];
    $book->typeOfName = $obj['typeOfName'];
    */
    
    }else{
        http_response_code(403);    
    }
    
}else{
    http_response_code(501);
    exit();   
}



?>