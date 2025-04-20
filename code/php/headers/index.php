<html>
  <head>
    <title>PHP</title>
  </head>
  <body>
  <?php
    foreach (getallheaders() as $name => $value) {
      echo "$name: $value\n";
    }
  ?>
  </body>
</html>
