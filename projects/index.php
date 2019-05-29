<?php
    global $title;
    $title = "Projects";
    require_once $_SERVER["DOCUMENT_ROOT"] . "/header.php";
?>
        <main>
            <h1><?php echo $title; ?></h1>
            <p>Miscellaneous stuff of mine that somehow ended up online.</p>
			<ul>
				<li><a href="https://na-und-goethe-2017.itch.io/na-und">Na Und?</a> - An experimental game of choice made during the <a href="https://www.goethe.de/en/uun/ver/arg/ort/bst.html">Goethe-Institut Boston</a>’s iteration of the 48-hour <a href="https://www.goethe.de/en/uun/ver/arg.html">Art Games</a> game jam.</li>
                <li><a href="/generate">Generate</a> - An intersection of computer science and art.</li>
				<li><a href="/projects/chimerical">Chimerical Colors</a> - An exploration in colors that don’t exist! Stare at the X until the screen changes color. Explanation included in my <a href="/article/?id=146059450673">Colors are Complicated blog post</a>.</li>
                <li><a href="https://www.youtube.com/watch?v=G1UExKuc1hQ">The Mechanical Orchestra</a> - Floppy drives that play music, controlled using an Arduino UNO.
                <li><a href="https://github.com/juliencherry/whats-up">What’s Up?</a> - An in-progress tool I’m designing to help manage reminders and todos.
				<li><a href="/projects/lambda-iota-engma">Lambda Iota Engma</a> - A concept website for the Northeastern linguistics club.</li>
				<li><a href="/projects/driving-charts">Driving in the USA</a> - A data visualization.</li>
				<li><a href="/projects/playing-cards">Playing Cards</a> - A playing card concept for my Art Foundations course (Google Chrome only). </li>
			</ul>
        </main>
<?php
    require_once $_SERVER["DOCUMENT_ROOT"] . "/footer.php";
?>
