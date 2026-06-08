const clipboardy = require('clipboardy');
const fs = require('fs');
const path = require('path');

const HISTORY_FILE = 'history.json';

function loadHistory() {
    try {
        return JSON.parse(fs.readFileSync(HISTORY_FILE, 'utf8'));
    } catch (err) {
        return [];
    }
}

function saveHistory(history) {
    fs.writeFileSync(HISTORY_FILE, JSON.stringify(history, null, 2));
}

async function main() {
    let history = loadHistory();
    let lastContent = await clipboardy.read();
    console.log('Clipboard Manager запущен. Ctrl+C для выхода.');
    setInterval(async () => {
        const current = await clipboardy.read();
        if (current !== lastContent && current.trim() !== '') {
            history.push({
                timestamp: new Date().toISOString(),
                content: current
            });
            saveHistory(history);
            console.log(`[${new Date()}] Скопировано: ${current.substring(0, 50)}...`);
            lastContent = current;
        }
    }, 1000);
}

main();
