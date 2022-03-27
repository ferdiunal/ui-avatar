# UI Avatar


# Install

```cmd

    git clone git@github.com:ferdiunal/ui-avatar.git uiAvatar

    cd uiAvatar

    go get -u .

    go run main.go

```


# Parameters

## Name

Generate avatar with name params

```
    https://....com/?name=John Doe

```

## Background (hex color)

Set background color

```
    https://....com/?name=John Doe&background=F80000&color=fff

```

## Font Color (hex color)

Set font color

```
    https://....com/?name=John Doe&background=fff&color=F80000

```

## Format (svg/png)

You can set the format of the avatar to be created as svg or png. Default: png


```
    https://....com/?name=John Doe&background=fff&color=F80000&format=svg

```

## Bold (boolean)

Determines whether to use bold font. Currently it only works in svg. Default false


```
    https://....com/?name=John Doe&background=fff&color=F80000&format=svg&bold=true

```

## Rounded (boolean)

It is determined whether the avatar to be created will be a circle or not. Default false



```
    https://....com/?name=John Doe&background=fff&color=F80000&format=svg&bold=true

```


## Font Size (float, min:0.1, max: 1, default: 0.5)

Sets the font size between 0.1 and 1. Default: 0.5


```
    https://....com/?name=John Doe&background=fff&color=F80000&format=svg&bold=true&font-size=0.7

```

## Image Size (float, min:0.1, max: 1, default: 0.5)

Sets the image size between 25x25 and 1024x1024 Default: 512


```
    https://....com/?name=John Doe&background=fff&color=F80000&format=svg&bold=true&font-size=0.7&size=125

```

## Initial Character (int)

Length of the generated initials. Default: 2


```
    https://....com/?name=John Doe&background=fff&color=F80000&format=svg&bold=true&font-size=0.7&size=125&length=1

```


## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.