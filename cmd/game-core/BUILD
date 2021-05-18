load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "game-core",
    embed = [":game-core_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "game-core_lib",
    srcs = ["main.go"],
    importpath = "github.com/delosrogers/game-engine/cmd/game-core",
    visibility = ["//visibility:private"],
    deps = ["@com_github_gin_gonic_gin//:gin"],
)
