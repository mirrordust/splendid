defmodule Mountain.Repo do
  use Ecto.Repo,
    otp_app: :mountain,
    adapter: Ecto.Adapters.Postgres
end
