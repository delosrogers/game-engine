load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "gamecore",
    embed = [":gamecore_lib"],
    visibility = ["//visibility:public"],
    data = [
        "//cmd/gamecore:home.html",
    ],
)

go_library(
    name = "gamecore_lib",
    srcs = ["main.go"],
    importpath = "github.com/delosrogers/game-engine/cmd/gamecore",
    visibility = ["//visibility:private"],
    deps = [
        "//internal/core",
        "@com_github_gin_gonic_gin//:gin",
        "@com_github_google_uuid//:uuid",
    ],
)

filegroup(
    name = "webfiles",
    srcs = ["home.html"],
)
