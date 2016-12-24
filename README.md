<img align="right" src="https://raw.githubusercontent.com/srs/govatar-net/master/misc/logo.jpeg">

# Tiny Server in GO serving avatars

This is a tiny server (written in GO) that serves unique avatars based on a string-hash. This string-hash can
be anything but usually the email or username. The avatars is rendered using the amazing
[Govatar](https://github.com/o1egl/govatar) library.

## Installing

To install the server, use the following:

```
$ go get -u github.com/srs/govatar-net
```

You can also use the `Makefile` for building on OSX, Windows and Linux:

```
$ make build
```

## Docker

And if you are using Docker (see [repository](https://hub.docker.com/r/stenrs/govatar-net/)):

```
$ docker run -d -p 8000:8000 stenrs/govatar-net
```

Or if you are using `docker-compose`, then add the following entries to `docker-compose.yml`.

```yml
govatar:
  image: stenrs/goavatar-net
  ports:
    - "8000:8000"
```

## Using

When starting the server it will default start on port `8000`. You can change this by setting the OS environment variable `PORT`. The server exposes one endpoint that renders the avatar.

```
http://localhost:8000/{gender}/{hash}.{ext}?size={size}
```

Here's an explanation about the various path-parameters and query-parameters:

* `gender` - This alternates between two skins - `f` for female and `m` for male.
* `hash` - A string that represents something unique. For example e-mail or user-name.
* `ext` - Image format to output - either `png` or `jpeg`.
* `size` - The size of the image to output. This is optional.

A couple of examples:

```
http://localhost:8080/m/mr@bar.com
http://localhost:8080/f/miss@bar.com
http://localhost:8080/f/miss@bar.com?size=200
```

# Demo Server

I have set up a demo server on Heroku (https://govatar.herokuapp.com) that you
can try out. It's a free instance so it can be pretty slow.

* https://govatar.herokuapp.com/m/mr@bar.com
* https://govatar.herokuapp.com/f/miss@bar.com
* https://govatar.herokuapp.com/f/miss@bar.com?size=200


## License

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
