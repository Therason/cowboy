mod pane;

use std::io;
use std::str;
use std::process::Command;
use tui::Terminal;
use tui::backend::TermionBackend;
//use tui::text::Text;
use termion::raw::IntoRawMode;
//use tui::widgets::{List, ListItem, Block, Borders};
//use tui::widgets::Widget;
//use tui::layout::{Layout, Constraint, Direction};

use pane::Pane;

fn main() -> Result<(), io::Error> {
    let stdout = io::stdout().into_raw_mode()?;
    let backend = TermionBackend::new(stdout);
    let mut terminal = Terminal::new(backend)?;
    terminal.hide_cursor()?;
    terminal.clear()?;

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
        })?
    }
}
