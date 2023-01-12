## Cric4CWC23

* Aims to develop a cricket trading api. 
* Predict playing 11 for CWC 2023
* API: `teamIndia.playing11(match)`
* Opponent, Venue, Conditions based algo.

## Execute
- `go run main.go`

## Data
With due credits to https://cricsheet.org/ 
Data being used in `csv` format because csv loaders are cheap.
- ODIs Men - `odis_male_csv/*.csv`
- Players - `people.csv`


## Steps
### Roles 
- Get batter list
- Get allrounder list
- Get bowler list
- Get wicketkeeper list
- Get captain list
### Stats
- #runs
- #matches
- #win percentage
- #wickets
- #stumps
- #bowled
- #30/50/70/100
### Match
- 11
- Result
- Opponent
- Venue
- Man of the Match
- Toss
- Best Every Role