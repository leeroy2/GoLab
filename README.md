# Project Title

Coding Exercise: Parsing Metacritic’s www content

## Getting Started

The annexed *.py files located in this folder, are the requested deliverables of the excercise.

### Deliverables

* sonyParser.py (Main class with a method returning the top PS4 games)
* sonyParserApp.py (Flask application; provides the requested API)
* sonyParserAppTests.py (unit testing application)
* requirements.txt (Provides all the necessary libraries)


### Prerequisites

For Linux machines, it is required Python 2.7. Open a linux terminal and Use the following command in order to confirm the output:

```
python --version
```

In order to install pip, open a linux terminal and run the following commands:

```
apt-get update
apt-get -y install python-pip
```

Also, a file named requirementes contains all the required libraries of the project; Type the following command in order to install the libraries:

```
pip install -r requirements.txt
```

### Execution

In order to start the client that provides the basic API, please use the following command

```
python sonyParserApp.py
```

The aforementioned command will start the application on port 5000 (default)

Open a browser and navigate [here](http://localhost:5000)

or alternatively confirm that is running by opening a new terminal and type:

```
curl http://localhost:5000

Web App for retrieving top PS4 games 

```

In order to view the top games as per [url](http://www.metacritic.com/game/playstation-4) please navigate to [this](http://localhost:5000/games) url

To query a game as per its title browse to [http://localhost:5000/games/<title-of-game>](http://localhost:5000/games/<title-of-game>) where you have to replace the <title-of-game> with the desired game title

## Running the tests

Unit testing required the service up n running. Before testing make sure that the application is up n running

### Break down into end to end tests

The following e2e unittest tests a happy flow of the application. More specifically it tests:

```
http://localhost:5000/ if the response is HTTP OK 200

http://localhost:5000/games if this is dict or title and score keys are missing, or a specific string is missing (Celest)

http://localhost:5000/games/Celeste if response code is HTTP OK 200 or its type is dict. Also it checks if Celeste is in the response object.

Finally it tests a random [url](http://localhost:5000/games/Microsoft) and ensures that the response data contains the string "game not found"


```

## Authors

* **Lambros Batalas** - *Initial work* - [LinkedIn](https://www.linkedin.com/in/lambros-batalas-2b403965/)
