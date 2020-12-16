use tui::widgets::{List, ListItem, Block, Borders};
use tui::layout::{Constraint, Direction, Layout};

//Parent represents the parent direcotry, current is the current direcotry, and Child is the
//contents of whatever the cursor is highlighting
#[allow(dead_code)]
pub enum PaneType {
    Parent,
    Current,
    Child,
}

//Pane is used to represent a singular block in the terminal 
#[allow(dead_code)]
pub struct Pane<'a> {
    items: Vec<ListItem<'a>>,
    ptype: PaneType,
}

impl Pane<'_> {
    //Create a new pane given a &str, which is then split into a list 
    pub fn new(content: &str, ptype: PaneType) -> Pane {
        let mut items = Vec::new();
        for (_, f) in content.split_whitespace().enumerate() {
            items.push(ListItem::new(f));
        }
        Pane {
            items,
            ptype,
        }

    } 

    //Render the pane
    pub fn visualize(
        self,
        t: &mut tui::Frame<tui::backend::TermionBackend<termion::raw::RawTerminal<std::io::Stdout>>>,
    ){
        let chunks = Layout::default()
            .direction(Direction::Horizontal)
            .constraints(
                [
                    Constraint::Percentage(20),
                    Constraint::Percentage(50),
                    Constraint::Percentage(30),
                ]
                .as_ref(),
            )
            .split(t.size());

        //this seciton might be better suited for a window struct 
        let size = match self.ptype {
            PaneType::Parent => chunks[0],
            PaneType::Current => chunks[1],
            PaneType::Child => chunks[2],
        };
    
        let list = List::new(self.items)
            .block(Block::default().title("Cowboy").borders(Borders::ALL));
        t.render_widget(list, size);

    }
}
