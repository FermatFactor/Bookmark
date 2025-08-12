// DOM Elements
const accentPicker = document.getElementById("accentColor");
const bookmarksContainer = document.getElementById("bookmarksContainer");
const bookmarkForm = document.getElementById("bookmarkForm");
const bookmarkText = document.getElementById("bookmarkText");

// Load accent color from localStorage
document.addEventListener("DOMContentLoaded", () => {
    const savedColor = localStorage.getItem("accentColor");
    if (savedColor) {
        document.documentElement.style.setProperty("--accent-color", savedColor);
        accentPicker.value = savedColor;
    }

    loadBookmarks();
});

// Change accent color in real time
accentPicker.addEventListener("input", (e) => {
    const color = e.target.value;
    document.documentElement.style.setProperty("--accent-color", color);
    localStorage.setItem("accentColor", color);
});

// Save bookmark
bookmarkForm.addEventListener("submit", (e) => {
    e.preventDefault();

    const text = bookmarkText.value.trim();
    if (!text) return;

    const bookmarks = getStoredBookmarks();
    bookmarks.push({ text, date: new Date().toISOString() });
    localStorage.setItem("bookmarks", JSON.stringify(bookmarks));

    bookmarkText.value = "";
    renderBookmarks(bookmarks);
});

// Get bookmarks from storage
function getStoredBookmarks() {
    return JSON.parse(localStorage.getItem("bookmarks")) || [];
}

// Load bookmarks on page load
function loadBookmarks() {
    const bookmarks = getStoredBookmarks();
    renderBookmarks(bookmarks);
}

// Render bookmarks in grid
function renderBookmarks(bookmarks) {
    bookmarksContainer.innerHTML = "";

    bookmarks.forEach((bm, index) => {
        const card = document.createElement("div");
        card.className = "bookmark-card";

        card.innerHTML = `
            <p>${bm.text}</p>
            <small>${new Date(bm.date).toLocaleString()}</small>
            <button class="delete-btn" onclick="deleteBookmark(${index})">Delete</button>
        `;

        bookmarksContainer.appendChild(card);
    });
}

// Delete a bookmark
function deleteBookmark(index) {
    const bookmarks = getStoredBookmarks();
    bookmarks.splice(index, 1);
    localStorage.setItem("bookmarks", JSON.stringify(bookmarks));
    renderBookmarks(bookmarks);
}

// Logout placeholder (will later connect to backend)
function logout() {
    alert("Logout functionality will be connected to backend authentication.");
}
