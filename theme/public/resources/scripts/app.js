function fixDateTimeString(dateTimeString) {
	const regex = /([\d-]+) ([\d:]+) ([\d+-]{3})([\d]{2})/;
	const results = regex.exec(dateTimeString);

	const dateString = results[1];
	const timeString = results[2];
	const tzFront = results[3];
	const tzBack = results[4];

	return `${dateString}T${timeString}.000${tzFront}:${tzBack}`
}

document.addEventListener("DOMContentLoaded", ()=> {
	const dateElements = document.querySelectorAll(".date");
	for (let index = 0; index < dateElements.length; index++) {
		const dateElement = dateElements[index];
		let dateString = fixDateTimeString(dateElement.textContent);
		const date = new Date(dateString);
		dateElement.innerHTML = date.toLocaleString()
	}

	const currentPath = location.pathname;
	const anchors = document.querySelectorAll("a");
	for (let index = 0; index < anchors.length; index++) {
		const anchor = anchors[index];
		const href = anchor.getAttribute("href");

		const isCurrentPage = href === currentPath
		if (isCurrentPage) {
			const newSpan = document.createElement("span");
			newSpan.innerHTML = anchor.innerHTML;
			anchor.replaceWith(newSpan);
		}
	}
});
