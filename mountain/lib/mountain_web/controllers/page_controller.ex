defmodule MountainWeb.PageController do
  use MountainWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
