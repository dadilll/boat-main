document.addEventListener("DOMContentLoaded", () => {
    const xCoord = document.getElementById('xCoord');
    const yCoord = document.getElementById('yCoord');
    const fireButton = document.getElementById('fireButton');

    const yourBoard = document.getElementById('yourBoard');
    const enemyBoard = document.getElementById('enemyBoard');

    const socket = new WebSocket('ws://localhost:8080/ws');
    let playerID = null;

    socket.addEventListener('open', (event) => {
        console.log('WebSocket connection established');
        // Request playerID from server
        socket.send(JSON.stringify({ type: "init" }));
    });

    socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);
        
        if (message.type === "init") {
            playerID = message.playerID;
        } else {
            const gameState = message;
            if (gameState.playerBoards && gameState.enemyBoards && playerID !== null) {
                updateBoard(yourBoard, gameState.playerBoards[playerID]);
                updateBoard(enemyBoard, gameState.enemyBoards[playerID]);
            } else {
                console.error('Received game state does not contain playerBoards or enemyBoards');
            }
        }
    });

    fireButton.addEventListener('click', () => {
        const x = parseInt(xCoord.value);
        const y = parseInt(yCoord.value);
        const move = { type: "move", x: x, y: y, playerID: playerID };
        socket.send(JSON.stringify(move));
    });

    function updateBoard(boardElement, boardData) {
        boardElement.innerHTML = '';
        boardData.Cells.forEach((row, y) => {
            const tableRow = document.createElement('tr');
            row.forEach((cell, x) => {
                const tableCell = document.createElement('td');
                tableCell.className = getClassForCell(cell);
                tableRow.appendChild(tableCell);
            });
            boardElement.appendChild(tableRow);
        });
    }

    function getClassForCell(cell) {
        switch (cell) {
            case 0: return 'empty';
            case 1: return 'ship';
            case 2: return 'hit';
            case 3: return 'miss';
            default: return '';
        }
    }

    function populateSelectOptions(selectElement, maxValue) {
        for (let i = 0; i < maxValue; i++) {
            const option = document.createElement('option');
            option.value = i;
            option.text = i;
            selectElement.appendChild(option);
        }
    }

    populateSelectOptions(xCoord, 10);
    populateSelectOptions(yCoord, 10);
});
