mod pane;
mod util;
mod app;

use crate::util::event::Event;
use crate::util::event::Events;

use std::io;
use std::error::Error;
use tui::Terminal;
use tui::backend::TermionBackend;
use termion::raw::IntoRawMode;
use termion::event::Key;

fn main() -> Result<(), Box<dyn Error>> {
    let stdout = io::stdout().into_raw_mode()?;
    let backend = TermionBackend::new(stdout);
    let mut terminal = Terminal::new(backend)?;
    terminal.hide_cursor()?;
    terminal.clear()?;
    let events = Events::new();

    //main loop 
    loop {
        //rendering 
        app::draw(&mut terminal)?;

        //keystrokes
        if let Event::Input(key) = events.next()? {
            if key == Key::Char('q') {
                terminal.clear()?;
                break;
            }
        }

    }
    Ok(())
}
