# README

## About

This is a demonstration of integrating Wails with FrankenPHP.

I have minimal Go knowledge, and even less C experience, so this demo just gets you to the Laravel welcome screen, and nothing beyond that.

Only tried building on MacOS, adjust for other platforms accordingly.

To build do the following:

1. `cd` into the php folder, follow the compilation instructions on the (FrankenPHP website)[https://frankenphp.dev/docs/compile/]
2. When running the configure script don't forget to add a flag for the OpenSSl extension. On my system the complete configuration command looked like so:
    ```bash
        ./configure \
    --enable-embed=static \
    --enable-zts \
    --disable-zend-signals \
    --disable-opcache-jit \
    --enable-static \
    --enable-shared=no \
    --with-iconv=/opt/homebrew/opt/libiconv/ \
    --with-openssl=/opt/homebrew/opt/openssl/
    ```
3. Open `main.go`, and replace "/path/to/public/index.php" with an absolute path to any PHP app. This repo contains a sample Laravel application in the `root` folder. If you use that don't forget to run `composer install` in that directory before building anything.
4. Run the build command: `CGO_CFLAGS=$(php-config --includes) CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" wails build -debug`. I was only able to build the app with the `-debug` flag.
5. Open `build/bin/[your-platform]` and run the executable file.
