# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run github.com/aofiee666/OmiseWallet

### Go to http://localhost:9000/ and you'll see:

![alt text](https://github.com/aofiee/OmiseWallet/blob/master/public/img/loginScreen.png?raw=true)


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