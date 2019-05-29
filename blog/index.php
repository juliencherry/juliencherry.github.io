<?php
	error_reporting(E_ERROR);
	global $title;
	$title = "Blog";
	require_once $_SERVER["DOCUMENT_ROOT"] . "/header.php";
	include "../vendor/autoload.php";
?>

<main>
	<h1><?php echo $title; ?></h1>
	<p>I try to write occasionally&mdash;here’s a sampler of my recent posts.</p>

	<?php // iterate through posts

	// authenticate via API Key
	$client = new Tumblr\API\Client('a6kE6bVie91SOrUteuprDKFy6Y4BAN1461fLXt0DLlZdaeTCnI');

	// make the request
	$posts = $client->getBlogPosts('juliencherry.tumblr.com', array('tag' => 'website'));

	foreach($posts->posts as $post) {
		echo "<li>";
		$date = DateTime::createFromFormat("Y-m-d H:i:s T", $post->date);

		if ($post->type == "text") {
			echo "<a href='/article?id=" . $post->id . "'>" . $post->title . "</a>";
		} else if ($post->type == "photo") {
			echo "<a href='/article?id=" . $post->id . "'>" . "(Untitled)" . "</a>";
		} else {
			// do nothing ¯\_(ツ)_/¯
		}

		echo "</li>";
	}

	?>
</main>

<?php require $_SERVER["DOCUMENT_ROOT"] . "/footer.php"; ?>
