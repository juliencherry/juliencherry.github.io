<?php
    error_reporting(E_ERROR);

    $id;
    $tag;

    if (!empty($_GET)) {
        if (!empty($_GET["id"])) {
            $id = htmlspecialchars($_GET["id"]);
        } else if (!empty($_GET["tag"])) {
            $tag = htmlspecialchars($_GET["tag"]);
        } else {
            // do nothing
        }
    }

    global $title;
    $title = "Article";
    require_once $_SERVER["DOCUMENT_ROOT"] . "/header.php";
    include "../vendor/autoload.php";
?>

<main>
    <?php // print article

        // authenticate via API Key
        $client = new Tumblr\API\Client('a6kE6bVie91SOrUteuprDKFy6Y4BAN1461fLXt0DLlZdaeTCnI');

        if (isset($id)) {
            $posts = $client->getBlogPosts('juliencherry.tumblr.com', array('id' => $id));
        } else if (isset($tag)) {
            $posts = $client->getBlogPosts('juliencherry.tumblr.com', array('tag' => $tag));
        } else {
            $posts = $client->getBlogPosts('juliencherry.tumblr.com', array('tag' => "tag"));
        }

        foreach($posts->posts as $post) {
            $date = DateTime::createFromFormat("Y-m-d H:i:s T", $post->date);
            $dateAsHTML = "<time datetime=" . $date->format("Y-m-d"). ">" . $date->format("F jS, Y") . "</time>";

            echo "<article>";

            if ($post->type == "text") {
                echo "<h2>" . $post->title . "</h2>";
                echo $dateAsHTML;
                echo $post->body;
            } else if ($post->type == "photo") {
                $photo = $post->photos[0]->alt_sizes[1];
                echo '<iframe frameborder="0" scrolling="no" width="' . $photo->width . '" height="' . $photo->height . '" src="' . $photo->url . '"></iframe>';
                echo "<br />";
                echo $dateAsHTML;
                echo $post->caption;
            } else {
                echo "Couldn’t display article :/";
            }

            echo "</article>";
        }
    ?>
</main>

<?php require $_SERVER["DOCUMENT_ROOT"] . "/footer.php"; ?>
