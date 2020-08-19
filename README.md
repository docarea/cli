# Only temporary

this script is a temporary example
a better CLI client is planned and under development


Use this Script by exporting the followingEnvironment Variables:
 * DOCAREA_DOCUMENTATION_ID
 * DOCAREA_CLIENT_ID
 * DOCAREA_CLIENT_SECRET

Then the script can be executed with the argument of the path to the documentation

As Example:
```
./demo.bash Path/to/Index.html/Containing/folder

```

Or just simple use the following line
```

curl -s "https://raw.githubusercontent.com/docarea/cli/master/demo.bash" | DOCAREA_DOCUMENTATION_ID=${{ secrets.DOCAREA_DOCUMENTATION_ID }} DOCAREA_CLIENT_ID=${{ secrets.DOCAREA_CLIENT_ID }} DOCAREA_CLIENT_SECRET=${{ secrets.DOCAREA_CLIENT_SECRET }} bash -s "Path/to/Index.html/Containing/folder"

```
