#!/usr/bin/env python3

from __future__ import annotations
import re
import sys

# reference: https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
SEMVER_REGEX = r"^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$"


def removeprefix(text: str, prefix: str) -> str:
    if text.startswith(prefix):
        return text[len(prefix) :]
    else:
        return text


def get_semver(text: str, prefix: str) -> str | None:
    if prefix != "":
        text = removeprefix(text, prefix)

    r = re.compile(SEMVER_REGEX)
    result = r.match(text)

    if result is None:
        raise Exception(f"target is not semver string: '{text}'")

    return (
        text,
        result["major"],
        result["minor"],
        result["patch"],
        result["prerelease"],
        result["buildmetadata"],
    )


def main():
    semver_str: str = ""
    prefix: str = ""

    try:
        semver_str = sys.argv[1]
        prefix = sys.argv[2]
    except IndexError:
        pass

    semver, major, minor, patch, prerelease, build_metadata = get_semver(
        semver_str, prefix
    )

    if prerelease is None:
        prerelease = ""

    if build_metadata is None:
        build_metadata = ""

    print(f"::set-output name=version::{semver}")
    print(f"::set-output name=major::{major}")
    print(f"::set-output name=minor::{minor}")
    print(f"::set-output name=patch::{patch}")
    print(f"::set-output name=prerelease::{prerelease}")
    print(f"::set-output name=build-metadata::{build_metadata}")


if __name__ == "__main__":
    try:
        main()
    except Exception as err:
        print(f"::error::raised error: {err}")
        sys.exit(1)
