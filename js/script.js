'use strict'

$(document).ready(function() {
	$('main').flowtype({
		fontRatio: 30,
		maxFont: 25,
		minFont: 20
	})

	$('a').not('a[href^="https://"]')
		.not('a[href^="mailto:"]')
		.not('a[href^="/"]')
		.not('a[href^="#"]')
		.addClass('unsafe-link')
		.prepend('<i class="fas fa-unlock-alt fa-sm"></i>')
})
