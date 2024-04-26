look at documentation first: does it make sense

initial implementation
    study documentation'requirements & tests; discuss
    study implementation; discuss
        Q: any inefficiency jumps at you?
        write down ideas

measure after modifying: suite of unit tests for correctness, suite of benchmarks for performance
    write some (sub)-benchmarks
        b.ResetTimer
        for range b.N
    run benchmarks
        explain options
            -run ^$
            -bench .

commit each time we make a change, so we can go back to compare benchmark results

ideas?
    algo?
    strings.EqualFold?
    eliminate normalization of names

low-hanging algorithmic fruit: store guest in a set data structure
    currently O(nm) but could be: O(n)
    benchstat to compare execution time
        installation: go install golang.org/x/perf/cmd/benchstat@latest
        basic command
        additional options for go test
            -benchtime 3s (or Nx)
            -count 10
        idle machine
        explain source of reduce variance: hungry apps, thermal throttling
        don't run them in the cloud on a shared machine
            show variance in jub0bs/cors
    Later: show that, when compared to initial approach, hashing dominates exec time in case of maliciously long name

rule of thumb: identify and eliminate unnecessary allocations
    link to Bryan Boreham
    -benchmem
    but don't know where they come from
    pprof
        explain how it works
        different types of profiles
        -memprofile mem.out
        go tool pprof mem.out
            top
            web

=> avoid splitting csv; instead, munch at it
    re-run pprof
        results now show that allocations are due to strings.Join

=> don't build any accepted slice; instead, simply return csv if success

idea: exit early if seeing a guest for the second time
    => clone set to keep track of seen guests
    pprof memory profile: allocations are back
    
idea: keep track of seen guests using a lighter data structure: []bool
    pprof memory profile: allocations are back

use a sync.Bool of *[]bool to amortize allocs
    explain sync.Pool
        must not be copied
        not typesafe: onus is on you (type assertion)
        New
        elements must be identical in size and value: clear before putting back
        *[]bool, not []bool
            see staticcheck rule: https://staticcheck.io/docs/checks#SA6002
    pprof memory profile: allocations are gone

write BouncerCheck/maliciously_long_non-invited_name later
    show that long execution time

=> augment sorted set with maxLen
    reduces execution time of new benchmark case

new requirement: names in csv must be sorted (in addition to lowercase and unique)
    guests must enter in lexicographical order
    => get rid of most complexity

bounds check
    go build -gcflags="-d=ssa/check_bce/debug=1" ./party/party.go
    add a dummy access after updating commaPos to save one
        _ = s[commaPos+1:]
    most beneficial in loops: show an example
        or write a benchmark case with many guests that shows a difference
        perhaps a good use case for showing the importance of statistical testing

TODO: fix doc comments
    go over all commits and make sure they make sense
    tidy notes
    publish github repo?
    produce slides?
    understand party.test file output from
        go test -run ^$ -bench . -benchmem -benchtime 3s -memprofile mem.out ./party

Aftermath
    - https://github.com/rs/cors/blob/8d33ca4794eae9bcb270e306fd3e9b89cf07ec4c/cors.go#L355
    - https://github.com/rs/cors/issues/170
    - https://github.com/rs/cors/pull/171 (a month to merge!)
    - https://deps.dev/go/github.com%2Frs%2Fcors/v1.11.0/dependents
    - https://github.com/google/certificate-transparency-go
    - https://bughunters.google.com/about/rules/4928084514701312/patch-rewards-program-rules
    - show email
    