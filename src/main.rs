mod pane;
mod util;

use crate::util::event::Event;
use crate::util::event::Events;

use std::io;
use std::error::Error;
use std::str;
use std::process::Command;
use tui::Terminal;
use tui::backend::TermionBackend;
use termion::raw::IntoRawMode;
use termion::event::Key;

use pane::Pane;

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
        terminal.draw(|t| {
            //read directory            
            let ls_dir = Command::new("ls").output()
                .expect("Unable to read directory"); 
            let files = str::from_utf8(&ls_dir.stdout).unwrap();
            
            //draw main pane 
            let main_pane = Pane::new(files);
            main_pane.visualize(t);
        })?;

        if let Event::Input(key) = events.next()? {
            if key == Key::Char('q') {
                break;
            }
        }

    }
    Ok(())
}
