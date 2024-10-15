use reqwest;
use serde::Deserialize;
use std::error::Error;

#[derive(Deserialize)]
struct Response {
    message: String,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let response = reqwest::get("http://localhost:8080/api/greet")
        .await?
        .json::<Response>()
        .await?;

    println!("Message from Go server: {}", response.message);
    Ok(())
}
