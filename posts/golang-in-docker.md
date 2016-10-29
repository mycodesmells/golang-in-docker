# Golang In Docker In 5 Minutes

If you want to deploy your application to the cloud, you may be considering to put it into a Docker container, so that you make sure the environment will be identical both in the cloud and locally. Great thing is, dockerizing your Go application won't take more that five minutes!

### Golang vs (most of) the field

One of the biggest advantages of Go language is that it produces a single binary file that doesn't require any additional dependencies. In comparison, in order to run Node JS application, you need to install your dependencies (we are talking about hundreds of megabytes), which can take much time (and has to be repeated every time you make any change to the code).

### Golang + Alpine Linux

If you have the comfort of having a single file to be run on the Docker, you may also want to save some time (and space) by using a tiny image. Probably the smallest one is called [Alpine Linux](https://www.alpinelinux.org/) which is only 5MB in size.

It sounds perfect, but you need to keep one thing in mind: Alpine is using _musl_ library (it is responsible for providing C interfaces for basic services for Linux-based systems, you can read more in [their FAQ](https://www.musl-libc.org/faq.html)), you need to compile your Go applications in a certain way (Thanks to **blang**'s [Github repo](https://github.com/blang/golang-alpine-docker) for that):

    CGO_ENABLED=0 go build -a -installsuffix cgo -o outputBinaryName main.go

Our sample application will be a small web server that shows some greeting to the users:

    // main.go
    ...
    func main() {
        http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
            rw.Write([]byte("Hello!"))
        })

        log.Fatal(http.ListenAndServe(":3000", nil))
    }


### Automation is everything

One thing I hate the most is manually doing something over and over again. This is why you should **automatize** building and running your application in Docker. First, we need to start by creating an image with `Dockerfile`:

    // docker/Dockerfile
    FROM alpine:edge
    ADD app /app
    RUN chmod 700 /app
    CMD "/app"

It's as simple as it can get: you just add a binary file of the application, make it executable within the container and start the server (by running the exec file).

Building the app is also very simple:

    // docker/build.sh
    #!/bin/bash
    echo "Building golang binary"
    CGO_ENABLED=0 go build -a -installsuffix cgo -o app ..
    echo "Building Docker image"
    docker build -t golang-in-docker .
    echo "Removing binary"
    rm app


First, we compile the application (using that Alpine-specific command), build the image from Dockerfile, and remove the binary (as it will be included in the image by the time we reach this statement).

Last but not least, let's automate running the container as well:

    // docker/run.sh
    #!/bin/bash
    docker run -p 13000:3000 -d golang-in-docker

This is just a single command, but thanks to having it in a script, we don't have to remember about port mapping (and everything we might be needing later on).

### Taking it for a spin

Let's build it and run the container:

    $ ./build.sh
    Building golang binary
    Building Docker image
    Sending build context to Docker daemon  5.64 MB
    Step 1 : FROM alpine:edge
     ---> a1a3cae7a75e
    Step 2 : MAINTAINER Paweł Słomka <pslomka@pslomka.com>
     ---> Using cache
     ---> de2540f3aa90
    Step 3 : ADD app /app
     ---> Using cache
     ---> 815ffae3d6d2
    Step 4 : RUN chmod 700 /app
     ---> Using cache
     ---> d7033589866d
    Step 5 : CMD "/app"
     ---> Using cache
     ---> eba08c1ef778
    Successfully built eba08c1ef778
    Removing binary

    $ ./run.sh
    14181df7157ef70cb004bcce99aa685dd98674d828a3b9422a1a4b52cfb0105e

    $ curl http://localhost:13000
    Hello!

The complete code of this example is available [on Github]().
