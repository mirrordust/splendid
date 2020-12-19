# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
use Mix.Config

config :mountain,
  ecto_repos: [Mountain.Repo]

# Configures the endpoint
config :mountain, MountainWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "1RhtfJDOumXCNSh42BrH9Sx4HImnfStDkWE84f0wknesbD0Zf/M1LIBXuzYJh9Jw",
  render_errors: [view: MountainWeb.ErrorView, accepts: ~w(html json), layout: false],
  pubsub_server: Mountain.PubSub,
  live_view: [signing_salt: "1TuWyH4Y"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
