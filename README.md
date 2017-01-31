# Simple realtime chat

This project was created as part of an
<a href="https://www.udemy.com/realtime-apps-with-reactjs-golang-rethinkdb">online course</a>.
I wrote some notes on the course; they are located in
<a href="https://github.com/mpillar/learning/blob/master/engineering/courses/developing-realtime-web-applications.md">this repository</a>.

## Running the application

### Frontend

    cd js && npm install && npm start

### Backend

#### Configure database

Install RethinkDB, and run the REQL commands listed in the `db/setup.reql` file.

#### Run Go code

    cd go && make install && make start
