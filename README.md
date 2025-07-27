## Storage Query Tool for Gno.land

### Prerequisites

* `gnokey` is installed and in your `PATH`
* `gnoland` and `gnoweb` are running

### Overview
This program automates querying the storage and deposit usage for each realm
and package. It calculates the total and average values.

* Recursively fetches all realm and package paths from `/r/`  and `/p/`
* Runs `gnokey` queries for each
* Aggregates total and average `storage` and `deposit` across all realms

---
To run against local node

```bash
go run main.go
```

To check against remote network
```
go run main.go -web https://staging.gno.land -rpc https://rpc.staging.gno.land
```

---
To query storage deposit for individual contract uses gnokey instead:

```bash
gnokey query vm/qstorage --data "gno.land/r/demo/boards"
```

output

```
height: 0
data: storage: 89558, deposit: 8955800
```
---

Output: with genesis_txs.jsonl loaded

`go run main.go`

```
Path                                          Storage (byte) Deposit (ugnot)
gno.land/r/agherasie/forms                          8858     885800
gno.land/r/demo/art/gnoface                         4090     409000
gno.land/r/demo/art/millipede                       3126     312600
gno.land/r/demo/atomicswap                         17486    1748600
gno.land/r/demo/banktest                            6840     684000
gno.land/r/demo/bar20                               9250     925000
gno.land/r/demo/boards                            134683   13468300
gno.land/r/demo/btree_dao                          15845    1584500
gno.land/r/demo/counter                             2861     286100
gno.land/r/demo/deep/very/deep                      2550     255000
gno.land/r/demo/disperse                           18837    1883700
gno.land/r/demo/draftrealm                          1755     175500
gno.land/r/demo/echo                                1695     169500
gno.land/r/demo/emit                                1766     176600
gno.land/r/demo/foo1155                            35069    3506900
gno.land/r/demo/foo20                              18402    1840200
gno.land/r/demo/foo721                             47054    4705400
gno.land/r/demo/games/shifumi                       7920     792000
gno.land/r/demo/grc20factory                       25244    2524400
gno.land/r/demo/grc20reg                           14825    1482500
gno.land/r/demo/groups                             70220    7022000
gno.land/r/demo/keystore                           10257    1025700
gno.land/r/demo/math_eval                           2066     206600
gno.land/r/demo/memeland                            9665     966500
gno.land/r/demo/microblog                           8248     824800
gno.land/r/demo/mirror                              5096     509600
gno.land/r/demo/profile                            21121    2112100
gno.land/r/demo/releases_example                    8838     883800
gno.land/r/demo/tamagotchi                          8357     835700
gno.land/r/demo/tests                              30878    3087800
gno.land/r/demo/tests/crossrealm                   28918    2891800
gno.land/r/demo/tests/crossrealm_b                  9065     906500
gno.land/r/demo/tests/crossrealm_c                  3323     332300
gno.land/r/demo/tests/init                          2524     252400
gno.land/r/demo/tests/subtests                      4695     469500
gno.land/r/demo/tests/test20                        5647     564700
gno.land/r/demo/tests_foo                           2581     258100
gno.land/r/demo/todolist                            9260     926000
gno.land/r/demo/types                              20065    2006500
gno.land/r/demo/ui                                  1833     183300
gno.land/r/demo/userbook                           11921    1192100
gno.land/r/demo/wugnot                             14197    1419700
gno.land/r/devrels/config                           8144     814400
gno.land/r/docs/adder                               4617     461700
gno.land/r/docs/avl_pager                         203284   20328400
gno.land/r/docs/avl_pager_params                   35216    3521600
gno.land/r/docs/buttons                             3587     358700
gno.land/r/docs/complexargs                         6612     661200
gno.land/r/docs/events                              2507     250700
gno.land/r/docs/hello                               1705     170500
gno.land/r/docs/home                                1846     184600
gno.land/r/docs/img_embed                           1747     174700
gno.land/r/docs/markdown                            1884     188400
gno.land/r/docs/minisocial                          1759     175900
gno.land/r/docs/minisocial/v1                       7954     795400
gno.land/r/docs/minisocial/v2                      18361    1836100
gno.land/r/docs/moul_md                            11288    1128800
gno.land/r/docs/optional_render                     1735     173500
gno.land/r/docs/resolveusers                        1915     191500
gno.land/r/docs/routing                             5904     590400
gno.land/r/docs/safeobjects                         4703     470300
gno.land/r/docs/source                              1719     171900
gno.land/r/docs/txlink                              3532     353200
gno.land/r/gfanton/gnomaze                         31226    3122600
gno.land/r/gnoland/blog                            47022    4702200
gno.land/r/gnoland/coins                           15805    1580500
gno.land/r/gnoland/events                          30118    3011800
gno.land/r/gnoland/faucet                          15645    1564500
gno.land/r/gnoland/ghverify                        17914    1791400
gno.land/r/gnoland/home                             5866     586600
gno.land/r/gnoland/monit                            8541     854100
gno.land/r/gnoland/pages                          174214   17421400
gno.land/r/gnoland/users                            7272     727200
gno.land/r/gnoland/users/v1                        63719    6371900
gno.land/r/gnoland/valopers                        81107    8110700
gno.land/r/gnoland/valopers_proposal                7196     719600
gno.land/r/gov/dao                                 34031    3403100
gno.land/r/gov/dao/v3/impl                         57593    5759300
gno.land/r/gov/dao/v3/init                          2915     291500
gno.land/r/gov/dao/v3/loader                        2243     224300
gno.land/r/gov/dao/v3/memberstore                 124166   12416600
gno.land/r/gov/dao/v3/treasury                     25226    2522600
gno.land/r/gov/dao/v3/treasury/test                 1065     106500
gno.land/r/grepsuzette/home                         5166     516600
gno.land/r/jjoptimist/home                          9571     957100
gno.land/r/leon/config                             16979    1697900
gno.land/r/leon/home                               13802    1380200
gno.land/r/leon/hor                               128558   12855800
gno.land/r/manfred/home                             1710     171000
gno.land/r/mason/home                               4831     483100
gno.land/r/matijamarjanovic/home                   22895    2289500
gno.land/r/matijamarjanovic/tokenhub               42670    4267000
gno.land/r/miko/calculator                         15051    1505100
gno.land/r/morgan/chess                            67638    6763800
gno.land/r/morgan/guestbook                        13651    1365100
gno.land/r/morgan/home                              2539     253900
gno.land/r/moul/config                              9114     911400
gno.land/r/moul/microposts                          4780     478000
gno.land/r/n2p5/config                             15843    1584300
gno.land/r/n2p5/haystack                            5099     509900
gno.land/r/n2p5/home                                9402     940200
gno.land/r/n2p5/loci                                6990     699000
gno.land/r/nemanya/config                           8306     830600
gno.land/r/nemanya/home                            24125    2412500
gno.land/r/nt/commondao                            89939    8993900
gno.land/r/stefann/registry                         6911     691100
gno.land/r/sunspirit/home                           2191     219100
gno.land/r/sunspirit/md                             2033     203300
gno.land/r/sys/names                                7306     730600
gno.land/r/sys/params                              14327    1432700
gno.land/r/sys/rewards                              1023     102300
gno.land/r/sys/txfees                               2475     247500
gno.land/r/sys/users                              615903   61590300
gno.land/r/sys/validators/v2                       16307    1630700
gno.land/r/ursulovic/registry                       8444     844400
gno.land/r/x/benchmark/storage                     22820    2282000
gno.land/r/x/map_delete                             3947     394700
gno.land/r/x/nir1218_evaluation_proposal           38805    3880500
gno.land/r/x/skip_height_to_skip_time               1071     107100

========== Totals ==========
Paths counted: 119
Total storage: 3012056 byte
Total deposit: 301205600 ugnot
========== Averages ==========
Average storage: 25311 byte
Average deposit: 2531139 ugnot

Path                                          Storage (byte) Deposit (ugnot)
gno.land/p/aeddi/panictoerr                         4252     425200
gno.land/p/agherasie/forms                         39081    3908100
gno.land/p/archives/bank                           10921    1092100
gno.land/p/demo/acl                                11162    1116200
gno.land/p/demo/avl                                30302    3030200
gno.land/p/demo/avl/list                            9612     961200
gno.land/p/demo/avl/pager                           9269     926900
gno.land/p/demo/avl/rolist                          7572     757200
gno.land/p/demo/avl/rotree                         12909    1290900
gno.land/p/demo/avlhelpers                          3048     304800
gno.land/p/demo/bf                                  3361     336100
gno.land/p/demo/blog                               28089    2808900
gno.land/p/demo/btree                              47939    4793900
gno.land/p/demo/cford32                            25123    2512300
gno.land/p/demo/combinederr                         3211     321100
gno.land/p/demo/context                             6482     648200
gno.land/p/demo/diff                                3314     331400
gno.land/p/demo/dom                                 4228     422800
gno.land/p/demo/entropy                             6477     647700
gno.land/p/demo/flow                               25881    2588100
gno.land/p/demo/fqname                              3474     347400
gno.land/p/demo/gnode                               3265     326500
gno.land/p/demo/gnorkle/agent                       4401     440100
gno.land/p/demo/gnorkle/feed                        4738     473800
gno.land/p/demo/gnorkle/feeds/static               10230    1023000
gno.land/p/demo/gnorkle/gnorkle                    20894    2089400
gno.land/p/demo/gnorkle/ingester                    3065     306500
gno.land/p/demo/gnorkle/ingesters/single            4138     413800
gno.land/p/demo/gnorkle/message                     4561     456100
gno.land/p/demo/gnorkle/storage                     2301     230100
gno.land/p/demo/gnorkle/storage/simple              4747     474700
gno.land/p/demo/grc/exts                            1414     141400
gno.land/p/demo/grc/grc1155                        39147    3914700
gno.land/p/demo/grc/grc20                          36359    3635900
gno.land/p/demo/grc/grc721                         71345    7134500
gno.land/p/demo/grc/grc777                          4946     494600
gno.land/p/demo/groups                              6419     641900
gno.land/p/demo/int256                             46831    4683100
gno.land/p/demo/json                              272843   27284300
gno.land/p/demo/math_eval/int32                    20900    2090000
gno.land/p/demo/memeland                           12088    1208800
gno.land/p/demo/merkle                              8406     840600
gno.land/p/demo/microblog                          11473    1147300
gno.land/p/demo/mux                                11978    1197800
gno.land/p/demo/nestedpkg                           5712     571200
gno.land/p/demo/ownable                            11743    1174300
gno.land/p/demo/ownable/exts/authorizable          16518    1651800
gno.land/p/demo/pausable                            6136     613600
gno.land/p/demo/rat                                 4722     472200
gno.land/p/demo/releases                           10765    1076500
gno.land/p/demo/seqid                               5326     532600
gno.land/p/demo/stack                               4358     435800
gno.land/p/demo/subscription                        2179     217900
gno.land/p/demo/subscription/lifetime              12439    1243900
gno.land/p/demo/subscription/recurring             14774    1477400
gno.land/p/demo/svg                                28202    2820200
gno.land/p/demo/tamagotchi                         11287    1128700
gno.land/p/demo/tests                               9455     945500
gno.land/p/demo/tests/p_crossrealm                  3696     369600
gno.land/p/demo/tests/subtests                      3129     312900
gno.land/p/demo/testutils                           7896     789600
gno.land/p/demo/todolist                            7512     751200
gno.land/p/demo/uassert                            23741    2374100
gno.land/p/demo/ufmt                               27371    2737100
gno.land/p/demo/ui                                 19281    1928100
gno.land/p/demo/uint256                           131189   13118900
gno.land/p/demo/urequire                           13663    1366300
gno.land/p/demo/watchdog                            4045     404500
gno.land/p/jeronimoalbi/datasource                 13766    1376600
gno.land/p/jeronimoalbi/datastore                  65291    6529100
gno.land/p/jeronimoalbi/expect                     64260    6426000
gno.land/p/jeronimoalbi/pager                      14534    1453400
gno.land/p/leon/coinsort                            6897     689700
gno.land/p/leon/ctg                                 3816     381600
gno.land/p/leon/pkgerr                              4270     427000
gno.land/p/leon/svgbtn                              9770     977000
gno.land/p/lou/ascii                               11840    1184000
gno.land/p/lou/query                               13223    1322300
gno.land/p/mason/md                                 6556     655600
gno.land/p/matijamarjanovic/charts                 15510    1551000
gno.land/p/morgan/chess                            39536    3953600
gno.land/p/morgan/chess/glicko2                     9543     954300
gno.land/p/morgan/chess/zobrist                    37013    3701300
gno.land/p/moul/addrset                             6259     625900
gno.land/p/moul/authz                              26195    2619500
gno.land/p/moul/collection                         16766    1676600
gno.land/p/moul/cow                                31501    3150100
gno.land/p/moul/debug                               5371     537100
gno.land/p/moul/dynreplacer                         4668     466800
gno.land/p/moul/errs                                4574     457400
gno.land/p/moul/fifo                                7935     793500
gno.land/p/moul/fp                                 14227    1422700
gno.land/p/moul/helplink                            6811     681100
gno.land/p/moul/md                                 22972    2297200
gno.land/p/moul/mdtable                             2557     255700
gno.land/p/moul/memo                                6260     626000
gno.land/p/moul/once                                8556     855600
gno.land/p/moul/printfdebugging                     6939     693900
gno.land/p/moul/realmpath                           5571     557100
gno.land/p/moul/txlink                              8593     859300
gno.land/p/moul/typeutil                           10433    1043300
gno.land/p/moul/ulist                              16982    1698200
gno.land/p/moul/ulist/lplist                       13095    1309500
gno.land/p/moul/web25                               3866     386600
gno.land/p/moul/xmath                              39392    3939200
gno.land/p/n2p5/brainfuck                           2749     274900
gno.land/p/n2p5/chonk                               5322     532200
gno.land/p/n2p5/haystack                           10934    1093400
gno.land/p/n2p5/haystack/needle                     8590     859000
gno.land/p/n2p5/loci                                3499     349900
gno.land/p/n2p5/mgroup                             18675    1867500
gno.land/p/nt/commondao                            94164    9416400
gno.land/p/nt/poa                                   9161     916100
gno.land/p/nt/treasury                             42526    4252600
gno.land/p/oxtekgrinder/ownable2step               14894    1489400
gno.land/p/pol/levenshtein                          8355     835500
gno.land/p/sunspirit/md                            20935    2093500
gno.land/p/sunspirit/table                          6595     659500
gno.land/p/sys/genesis                              4238     423800
gno.land/p/sys/validators                           5069     506900
gno.land/p/thox/accesscontrol                      17714    1771400
gno.land/p/thox/snowflake                           7236     723600
gno.land/p/thox/timelock                           21607    2160700
gno.land/p/wyhaines/rand/isaac                     13099    1309900
gno.land/p/wyhaines/rand/isaac64                   15527    1552700
gno.land/p/wyhaines/rand/xorshift64star            10684    1068400
gno.land/p/wyhaines/rand/xorshiftr128plus          10865    1086500

========== Totals ==========
Paths counted: 127
Total storage: 2225121 byte
Total deposit: 222512100 ugnot
========== Averages ==========
Average storage: 17520 byte
Average deposit: 1752063 ugnot
```
