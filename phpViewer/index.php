<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <style>
    .tdw{

    }
    </style>
    <script>
        $(document).ready(function(){
            $("#inPageSearcher").on("keyup", function() {
                var value = $(this).val().toLowerCase();
                $("#roll tr").filter(function() {
                $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
                });
            });
            });

    </script>
    </head>
    <body>
<?php
$thead='  
<thead>
<tr>
  <th scope="col">#</th>
  <th scope="col">priority</th>
  <th scope="col">status</th>
  <th scope="col">topics</th>
  <th scope="col">timestamp</th>
  <th scope="col w-100">msg</th>
  
</tr>
</thead>';
ini_set('error_reporting', E_ALL);
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);


include_once("rb.php");

R::setup( 'mysql:host=localhost;dbname=syslogManager','syslog', 'password' ); //for both mysql or mariaDB
//R::exec('PRAGMA foreign_keys = ON');
//$query = R::findAll( 'mainrall' );
$query = R::getAll( 'SELECT * FROM `mainrall` ORDER BY `mainrall`.`id` DESC ' ); 
// а можно и так  $query = R::getAll( 'SELECT * FROM jobs' ); 
echo '<pre>';
#var_dump($query[0]);
function priorityColor($index){
$style="";
    if($index==7){
        $style='table-primary';
    }
    if($index==6){
        $style='table-primary';
    }
    if($index==5){
        $style='table-warning';
    }
    if($index=='4'){
        $style='table-danger';
    }
    if($index==3){
        $style='bg-danger';
    }
    if($index==2){
        $style='bg-danger';
    }
    if($index==1){
        $style='bg-danger';
    }
    if($index==0){
        $style='bg-danger';
    }
    return $style;
}
echo'
<input class="form-control" id="inPageSearcher" type="text" placeholder="Искать">
<table class="table table-sm table-dark">'.$thead.'
<tbody id="roll">';
foreach( $query as $i){
    //var_dump($i);
    echo'<tr>';
    echo'<th scope="row">'.$i['id'].'</th>';
    echo'<td>'.$i['priority'].'</td>';
    echo'<td class='.priorityColor($i['priority']).'></td>';
    echo'<td>'.$i['topics'].'</td>';
    
    echo'<td>'.$i['timestamp'].'</td>';
    echo'<td>'.$i['msg'].'</td>';
    echo'</tr>';
    
}
echo '</tbody></table>';
//var_dump($query);
echo '</pre>';
//$book = R::dispense( 'mainrall' );
//$book->raw = 'ЕБаааааать!';
//$book->priority = 999;
//R::store( $book );
?>

</body></html>