import config from "./config.js";
var board;

var PostBoard = async function (e) {
    Swal.fire({
        imageUrl: config.url+"/"+config.loadingImg,
        title: "Checking board...",
        allowOutsideClick: false,
        allowEscapeKey: false,
        allowEnterKey: false,
        showConfirmButton: false,
    });

    var element = e.currentTarget
    var className = element.className;
    var newClassName = className.replace(/unscratched/g, "scratched");
    element.className = newClassName;

    var id = element.id;

    var idEx = id.split("-");
    var colArr = idEx[0].split("_");
    var rowArr = idEx[1].split("_");

    var id = board.id
    var col = colArr[1];
    var row = rowArr[1];

    var data = {
        id,
        col,
        row
    }

    var fullUrl = new URL(config.apiUrl + "/boards");

    var response = await fetch(fullUrl, {
        method: "POST", // *GET, POST, PUT, DELETE, etc.
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({...data})
    });
    var json = await response.json();

    var resData = json.data
    if (resData.state) {
        board = resData
        showPrize()
    } else {
        var img = element.children[0]
        img.src = config.url+'/'+resData.path
        element.removeEventListener('click', PostBoard)
        Swal.close()
    }
};

window.getBoard = async function getBoard() {
    Swal.fire({
        imageUrl: config.url+"/"+config.loadingImg,
        title: "Fetching board...",
        allowOutsideClick: false,
        allowEscapeKey: false,
        allowEnterKey: false,
        showConfirmButton: false,
    });

    var key = window.document.getElementById("key");
    var button = window.document.getElementById("button");

    var id = key.value;
    var fullUrl = new URL(config.apiUrl + "/boards");
    fullUrl.search = new URLSearchParams({ id }).toString();
    var response = await fetch(fullUrl);
    var json = await response.json();
    board = json.data;

    generatePrize()

    var masterBoard = window.document.getElementById("master_board");
    masterBoard.innerHTML = "";

    board.state.forEach((element, col) => {
        element.forEach((element2, row) => {
            if (element2 != -1) {
                var prize = board.prize[element2];
                generatePiece(col, row, prize.path, true);
            } else {
                generatePiece(col, row, config.defaultImg, false);
            }
        });
    });

    key.setAttribute("disabled", true);
    button.className =
        "bg-gray-400 px-4 rounded-lg border max-w-xs mx-6 max-h-9 text-white flex-none cursor-not-allowed";
    button.onclick = null;

    if (board.obtained_prize != -1) {
        showPrize()
    } else {
        setTimeout(Swal.close(), 5000);
    }

}

function generatePrize() {
    var dom = window.document.getElementById("master_prize")
    dom.innerHTML = ""

    board.prize.forEach(data => {
        var prizeElem = window.document.createElement("div")
        prizeElem.className = "flex justify-between p-6 my-2 max-w-sm  rounded-lg border shadow-md hover:bg-gray-700 bg-gray-800 border-gray-700"

        var img = window.document.createElement("img")
        img.src = config.url+'/'+data.path
        img.className = "max-h-6 border-white"

        var text = window.document.createElement("div")
        text.className = "px-1 text-white"
        text.innerHTML = data.name

        prizeElem.appendChild(img)
        prizeElem.appendChild(text)
        dom.appendChild(prizeElem)
    })

}

function printBoard() {
    console.log("PepeLaugh");
    console.log(board);
}

function generatePiece(col, row, src, scratch) {
    let ncol = `col_${col}`;
    let nrow = `row_${row}`;
    var dom = window.document.createElement("div");

    var className =
        "flex justify-center p-6 max-w-sm rounded-lg border shadow-md hover:bg-gray-700 bg-gray-800 border-gray-700";
    if (scratch) {
        className = className + " scratched";
    } else {
        className = className + " unscratched";
    }
    dom.className = className;
    dom.id = `${ncol}-${nrow}`;

    var img = window.document.createElement("img");
    img.className = "h-24";
    img.src = src;

    dom.appendChild(img);
    dom.addEventListener('click', PostBoard)
    window.document.getElementById("master_board").appendChild(dom);
}

function showPrize() {
    var prize = board.prize[board.obtained_prize]
    Swal.fire({
        imageUrl: config.url+'/'+prize.path,
        title: "Congratulation!",
        text: "You got " + prize.name,
        allowOutsideClick: false,
        allowEscapeKey: false,
        allowEnterKey: false,
        showConfirmButton: false,
    });
    cleanup()
}

function cleanup() {
    const elements = window.document.getElementsByClassName("unscratched");
    for (var i = 0; i < elements.length; i++) {
        elements[i].removeEventListener('click', PostBoard);
    }
}