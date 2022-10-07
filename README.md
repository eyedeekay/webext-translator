WebExtension Translator
=======================

This is a small application for generating translations from a `messages.json` file
used by a WebExtension. It loads a `messages.json` file, gathers tranlatable strings
from it, passes them to a LibreTranslate instance(my libretranslate instance within
I2P), and uses it to generate a new `messages.json` file containing the translations.

It is designed to be run from the root of the WebExtension directory where the
`_locales/` directory is located.

Installation
------------

Install it with:

```sh
go install github.com/eyedeekay/webext-translator@latest
```

then, to run it:

```sh
~/go/bin/webext-translator
```

Usage
-----

It is possible to generate one translation at a time, or all translations at once.

Russian-to-French example:

```sh
webext-translator -base="_locales/ru/messages.json" -lang="fr"
```

Frensh-to-Arabic-Chinese-English-German-Italian-Japanese-Portuguese-Russian-Spanish example

```sh
webext-translator -base="_locales/fr/messages.json"
``

Because it needs a default, it will automatically look for an English translation in
`_locales/en/messages.json` if no arguments are passed, and will generate translations
for all supported languages.
