# Card game paradise
Welcome to the world of card games. Card games are an amazing way to pass your time and, at the same time, strategize when playing games like poker, 21, or blackjack. The purpose of this repository is to enable robust backend API support that allows you to play any card game with ease.

### How to run
You can run this project on your local machine provided that you meet the prerequisites required for this application.

##### Prerequisite
1. You should have Golang installed on your system. If already installed, you can proceed ahead; otherwise, download it using the link [here](https://go.dev/doc/install).
2. Postman is required for testing APIs and executing BDD tests. If already installed, you can proceed ahead; otherwise, download it using the link [here](https://www.postman.com/downloads/).

##### Running application locally
1. Clone this repository using the command `git clone https://github.com/varadekd/card-game.git`.
2. Enter the folder using the command `cd card-game`.
3. Export the default path for the file. Please ensure you're in the root directory when exporting. You can use the command `export DEFAULT_CARDS_FILE_STORAGE=<working_dir>/card-game/data/cards.json`.
4. Start the application by running the command `go run main.go`. Please ensure you are in the root directory when starting the application.

The application will start on port 8080. You can import them using this [link](https://api.postman.com/collections/468401-0a3dbf26-2d93-4468-930a-cef0268f1c8d?access_key=PMAT-01HKF6R4XE016MVHWDZAWNZVQ2).

##### Running TDD Tests Locally
1. Export the default path for the file. Please ensure you're in the root directory when exporting. You can use the command `export DEFAULT_CARDS_FILE_STORAGE=<working_dir>/card-game/data/cards.json`.
2. To start the test execution, run the command `go test ./tests/...`. Please ensure you're in the root directory when executing the test cases.


Note: All the TDD test cases are located under the directory called "tests."

##### Running BDD Tests Locally
1.  Import this [link](https://api.postman.com/collections/468401-aedab6da-fc7c-4ad3-8868-01a0d8ac7149?access_key=PMAT-01HKF693R7FWYC11HNHMTXY856) into Postman.
2.  Right-click on the imported collection and click on Run Collection.
3.  You will now see all the APIs that you want to run. Click on the button labeled Run.
4.  Once the execution is completed, you will get the list of all the test cases run and their results.

### Tech used
- We used **Go** for the development of this project.
- We used **Gin-Gonic** for providing HTTPS functionality.
- We used **chai.js** for writing BDD test cases

### Development machine and tech versions
The APIs were developed and tested on Ubuntu 22.04 LTS. In case you are facing any troubles, please write back to varadekushagra@gmail.com.