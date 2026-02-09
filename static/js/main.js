// ============================================
// Groupie Tracker - JavaScript Interactions
// ============================================
// Navbar scroll effect
(function () {
    const navbar = document.getElementById('navbar');
    if (!navbar) return;
    window.addEventListener('scroll', function () {
        if (window.scrollY > 20) {
            navbar.classList.add('scrolled');
        } else {
            navbar.classList.remove('scrolled');
        }
    });
})();
// Mobile menu toggle
function toggleMobileMenu() {
    const menu = document.getElementById('mobile-menu');
    if (menu) {
        menu.classList.toggle('open');
    }
}
// Filter panel toggle
function toggleFilters() {
    const panel = document.getElementById('filter-panel');
    const chevron = document.getElementById('filter-chevron');
    if (panel) {
        panel.classList.toggle('open');
        if (chevron) {
            chevron.style.transform = panel.classList.contains('open') ? 'rotate(180deg)' : 'rotate(0)';
        }
    }
}
// Search form — you can hook this into your Go backend
// For client-side filtering or form submission:
(function () {
    const searchInput = document.getElementById('search-input');
    if (!searchInput) return;
    // Option 1: Submit as form (uncomment and wrap in a <form>)
    // searchInput.closest('form')?.addEventListener('submit', function(e) { ... });
    // Option 2: Client-side filtering (basic example)
    searchInput.addEventListener('input', function () {
        const query = this.value.toLowerCase();
        const cards = document.querySelectorAll('.artist-card');
        cards.forEach(function (card) {
            const name = card.querySelector('h3')?.textContent?.toLowerCase() || '';
            card.style.display = name.includes(query) ? '' : 'none';
        });
    });
})();