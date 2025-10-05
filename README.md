# Resenum

A small utility to enumerate possible resolutions following a given ratio.

## Usage

```
$ resenum --help
Usage: resenum [--min=] [--max=] [--standard] [--help] {ratio}
  -help
        Show help
  -max string
        Maximum resolution, expressed as X:Y. Each dimension is optional (default "1920:1080")
  -min string
        Minimum resolution, expressed as X:Y. Each dimension is optional (default "0:0")
  -standard
        Standard resolutions only, i.e. SD, HD, FHD, QHD and UHD
Examples:
        resenum 16:9
        resenum 16:9 --min=256:
        resenum 16:9 --max=:1080
```

### Sample output

```
$ resenum --standard 16:9
+------------+---------+
| RESOLUTION | NOTE    |
+------------+---------+
| 640x360    | SD      |
| 1280x720   | HD      |
| 1920x1080  | Full HD |
+------------+---------+
| TOTAL      | 3       |
+------------+---------+
```
