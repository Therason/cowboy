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
use pane::PaneType;

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
            //redering probably shouldn't be handled by main, so the need for a window struct has
            //emerged that handles rendering the panes as well as setting the size of each pane 

            //read parent directory
            let parent_ls = Command::new("ls").output()
                .expect("Unable to read directory");
            let parent_files = str::from_utf8(&parent_ls.stdout).unwrap();

            //draw parent pane 
            let parent_pane = Pane::new(parent_files, PaneType::Parent);
            parent_pane.visualize(t);

            //read current directory            
            let curr_ls = Command::new("ls").output()
                .expect("Unable to read directory"); 
            let curr_files = str::from_utf8(&curr_ls.stdout).unwrap();
            
            //draw main pane 
            let main_pane = Pane::new(curr_files, PaneType::Current);
            main_pane.visualize(t);
            
            //read child directory
            let child_ls = Command::new("ls").output()
                .expect("Unable to read directory");
            let child_files = str::from_utf8(&child_ls.stdout).unwrap();

            //draw child pane 
            let child_pane = Pane::new(child_files, PaneType::Child);
            child_pane.visualize(t);
        })?;

        if let Event::Input(key) = events.next()? {
            if key == Key::Char('q') {
                terminal.clear()?;
                break;
            }
        }

    }
    Ok(())
}
