# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Reproducible Installations
Install the dependencies and revisions listed in the lock file into the vendor directory. If no lock file exists an update is run.

```
$ glide install
```

or

```
docker run --rm -it -v $(pwd):/go/src/omise-go -w /go/src/omise-go instrumentisto/glide install
```

### Start the web server:

```
docker-compose build
docker-compose up -d
```

### Go to http://localhost:9000/ and you'll see:

| Username     | Password    |
| -------------|-------------|
| admin        | password    |

![alt text](https://OmiseWallet/blob/master/public/img/loginScreen.png?raw=true)


## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## UnitTest

http://localhost:9000/@tests