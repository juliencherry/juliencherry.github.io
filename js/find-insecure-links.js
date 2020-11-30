'use strict'

$(document).ready(function() {
	$('a').not('a[href^="https://"]')
		.not('a[href^="mailto:"]')
		.not('a[href^="/"]')
		.not('a[href^="#"]')
		.addClass('insecure-link')
		.prepend('<i class="fas fa-unlock-alt fa-sm"></i>')
})
